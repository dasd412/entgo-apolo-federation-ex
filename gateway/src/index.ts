import { ApolloServer } from "@apollo/server";
import { ApolloGateway, IntrospectAndCompose, RemoteGraphQLDataSource } from "@apollo/gateway";
import jwt from "jsonwebtoken";
import express from "express";
import cors from "cors";
import bodyParser from "body-parser";
import { expressMiddleware } from "@apollo/server/express4";
import { parse } from "graphql";

//todo console.log와 jwt 시크릿 전부 실무 환경에선 제거하거나 교체해야 함

// ✅ Auth 서버와 동일하게 유지
const JWT_SECRET = "your_secret_key";

// 🔹 **인증이 필요 없는 GraphQL Mutation 및 Query 목록**
const publicMutations = new Set(["signup", "login"]);
const publicQueries = new Set(["__schema", "IntrospectionQuery"]);

// 🔹 **각 마이크로서비스에 `userId`, `role`을 전달하는 DataSource**
class AuthenticatedDataSource extends RemoteGraphQLDataSource {
    willSendRequest({ request, context }) {
        if (context.userId) {
            request.http.headers.set("user-id", context.userId);
        }
        if (context.role) {
            request.http.headers.set("role", context.role);
        }
    }
}

// 🔹 **Apollo Gateway 설정**
const gateway = new ApolloGateway({
    supergraphSdl: new IntrospectAndCompose({//todo production에선 권장하지 않음.
        subgraphs: [
            { name: "user", url: "http://localhost:8081/graphql" },
            { name: "order", url: "http://localhost:8082/graphql" },
            { name: "delivery", url: "http://localhost:8083/graphql" },
        ],
    }),

    buildService({ name, url }) {
        return new AuthenticatedDataSource({ url });
    },
});

// 🔹 **JWT 검증 후 `userId` 및 `role`을 `context`에 저장**
async function getUserFromToken(token: string) {
    if (!token) return null;

    try {
        // ✅ HS256 알고리즘 사용 & Auth 서버와 동일한 시크릿 키 사용
        const decoded: any = jwt.verify(token.replace("Bearer ", ""), JWT_SECRET, { algorithms: ["HS256"] });

        return { userId: decoded.sub, role: decoded.role };
    } catch (error) {
        console.error("❌ Invalid token", error);
        return null;
    }
}

// 🔹 **GraphQL 요청에서 mutation 또는 query 필드명을 추출하는 함수**
function getOperationFieldName(query: string): { type: string; name: string | null } {
    try {
        const ast = parse(query); // GraphQL AST(Abstract Syntax Tree) 파싱
        const firstOperation = ast.definitions.find((def) => def.kind === "OperationDefinition");

        if (firstOperation?.kind === "OperationDefinition") {
            const operationType = firstOperation.operation; // `query` 또는 `mutation`
            const firstField = firstOperation.selectionSet.selections[0]; // 첫 번째 필드명 가져오기

            if ("name" in firstField) {
                return { type: operationType, name: firstField.name.value };
            }
        }
    } catch (error) {
        console.error("❌ Failed to parse GraphQL query:", error);
    }
    return { type: "unknown", name: null };
}

// 🔹 **Express 기반 Apollo Server 실행**
async function startServer() {
    const server = new ApolloServer({ gateway });
    await server.start();

    const app = express();

    // ✅ `body-parser` 추가하여 `req.body` 접근 가능하게 설정
    app.use(cors());
    app.use(bodyParser.json());

    // ✅ Apollo GraphQL 미들웨어 추가 (`context` 설정)
    app.use(
        "/graphql",
        expressMiddleware(server, {
            context: async ({ req }) => {
                try {
                    // 🔥 `req.body.query`에서 operationType (query/mutation) 및 필드명 추출
                    const { type, name } = req.body?.query
                        ? getOperationFieldName(req.body.query)
                        : { type: "unknown", name: null };

                    console.log(`🚀 GraphQL Request - Type: ${type}, Name: ${name}`);

                    // 🔹 인증이 필요 없는 Mutation 또는 Query는 JWT 검증 없이 통과
                    if ((type === "mutation" && name && publicMutations.has(name)) ||
                        (type === "query" && name && publicQueries.has(name))) {
                        return {};
                    }

                    // 🔹 JWT 검증 수행
                    const token = req.headers.authorization || "";
                    console.log("🔍 Received Token:", token);
                    const user = await getUserFromToken(token);

                    if (!user) {
                        throw new Error("Unauthorized");
                    }

                    return user;
                } catch (error) {
                    console.error("❌ Authentication error:", error);
                    throw new Error("Unauthorized");
                }
            },
        })
    );

    app.listen(4000, () => {
        console.log(`🚀 Apollo Gateway ready at http://localhost:4000/graphql`);
    });
}

startServer();
