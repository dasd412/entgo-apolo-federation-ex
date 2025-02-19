package auth

import (
	"context"
	"entgo.io/ent/privacy"
)

func DenyIfNoAuthority() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		role := UserAuthorityFromContext(ctx)

		if role == nil {
			return privacy.Denyf("no role found in context")
		}
		return privacy.Skip
	})
}

func AllowIfSignupOrLogin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		apiOperationName, err := ApiOperationNameFromContext(ctx)

		if err != nil {
			return privacy.Deny
		}

		if apiOperationName == "signup" || apiOperationName == "login" {
			return privacy.Allow
		}

		return privacy.Skip
	})
}

func AllowIfAdminOrAuthor() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		role := UserAuthorityFromContext(ctx)

		if role.IsAdmin() || role.IsAuthor() {
			return privacy.Allow
		}

		return privacy.Deny
	})
}
