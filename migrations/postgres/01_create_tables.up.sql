CREATE TABLE "category" (
  "id" uuid PRIMARY KEY,
  "name" varchar not null,
  "created_at" timestamp default current_timestamp not null,
  "updated_at" timestamp
);

CREATE TABLE "client" (
  "id" uuid PRIMARY KEY,
  "first_name" varchar not null,
  "last_name" varchar not null,
  "phone_number" varchar NOT NULL,
  "created_at" timestamp default current_timestamp not null,
  "updated_at" timestamp
);

CREATE TABLE "product" (
  "id" uuid PRIMARY KEY,
  "name" varchar,
  "category_id" uuid not null REFERENCES "category" ("id"),
  "description" varchar,
  "price" float not null,
  "quantity" integer,
  "created_at" timestamp default current_timestamp,
  "updated_at" timestamp
);

CREATE TABLE "order" (
  "id" uuid PRIMARY KEY,
  "client_id" uuid not null REFERENCES "client" ("id"),
  "price" float,
  "status" varchar default 'new',
  "created_at" timestamp default current_timestamp,
  "updated_at" timestamp
);

CREATE TABLE "order_products" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
  "order_id" UUID NOT NULL REFERENCES "order" ("id"),
  "product_id" UUID NOT NULL REFERENCES "product" ("id")
);