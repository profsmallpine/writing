package migrations

import "github.com/xy-planning-network/trails/postgres"

var List = []postgres.Migration{
	{Executor: createArticlesTable, Key: "202211220905_create_articles"},
}
