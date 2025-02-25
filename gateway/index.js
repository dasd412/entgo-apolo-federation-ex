const { ApolloGateway, IntrospectAndCompose } = require("@apollo/gateway");
const { ApolloServer } = require("@apollo/server");
const { expressMiddleware } = require("@apollo/server/express4");
const express = require("express");
const http = require("http");
const { WebSocketServer } = require("ws");
const { useServer } = require("graphql-ws/lib/use/ws");

const gateway = new ApolloGateway({
    supergraphSdl: new IntrospectAndCompose({
        subgraphs: [
            { name: "user", url: "http://localhost:8081/graphql" },
            { name: "order", url: "http://localhost:8082/graphql" },
            { name: "delivery", url: "http://localhost:8083/graphql" },
        ],
    }),
});

async function startServer() {
    // 🔹 Express + HTTP 서버 생성
    const app = express();
    const httpServer = http.createServer(app);

    // ✅ [NEW] JSON 요청을 처리할 수 있도록 미들웨어 추가
    app.use(express.json());

    // ✅ [NEW] 기본 경로(/) 처리: GraphQL Playground로 안내
    app.get("/", (req, res) => {
        res.send(`
      <h1>🚀 Apollo Gateway</h1>
      <p>Go to <a href="/graphql">GraphQL Playground</a></p>
    `);
    });

    // 🔹 Apollo Server 생성
    const server = new ApolloServer({ gateway });

    await server.start();
    app.use("/graphql", expressMiddleware(server));

    // 🔹 WebSocket Server 설정 (Subscription 지원)
    const wsServer = new WebSocketServer({
        server: httpServer, // ✅ HTTP 서버와 연결
        path: "/graphql",
    });

    useServer({ schema: gateway.schema }, wsServer);

    // 🔹 HTTP 서버 실행
    httpServer.listen(4000, () => {
        console.log(`🚀 Apollo Gateway running at http://localhost:4000/graphql`);
        console.log(`🚀 WebSocket Server running at ws://localhost:4000/graphql`);
    });
}

startServer();
