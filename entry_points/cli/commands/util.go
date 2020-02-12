package commands

import "lendme/golang-seeder-backend/src/di"

type DiAwareCommand func(di.Di, []string) string