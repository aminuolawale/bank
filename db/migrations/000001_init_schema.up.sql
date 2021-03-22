CREATE TABLE "accounts" (
    "id" bigserial primary key,
    "owner" VARCHAR NOT NULL,
    "balance" bigint NOT NULL,
    "currency" VARCHAR NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()));
CREATE TABLE "entries" (
        "id" bigserial primary key,
        "account_id" bigint NOT NULL,
        "amount" bigint NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now()));
CREATE TABLE "transfers" (
            "id" bigserial primary key,
            "from_account_id" bigint NOT NULL,
            "to_account_id" bigint NOT NULL,
            "amount" bigint NOT NULL,
            "created_at" timestamptz NOT NULL DEFAULT (now()));
            ALTER TABLE
                "entries"
            ADD
                foreign key ("account_id") references "accounts" ("id");
            ALTER TABLE
                "transfers"
            ADD
                foreign key ("from_account_id") references "accounts" ("id");
            ALTER TABLE
                "transfers"
            ADD
                foreign key ("to_account_id") references "accounts" ("id");
CREATE INDEX
                ON "accounts" ("owner");
CREATE INDEX
                ON "entries" ("account_id");
CREATE INDEX
                ON "transfers" ("from_account_id");
CREATE INDEX
                ON "transfers" ("to_account_id");
CREATE INDEX
                ON "transfers" (
                    "from_account_id",
                    "to_account_id"
                );
COMMENT
                ON column "entries"."amount" IS 'can be negative or positive';
COMMENT
                ON column "transfers"."amount" IS 'must be positive';
