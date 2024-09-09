DEFINE TABLE projects SCHEMAFULL;

DEFINE FIELD id ON projects TYPE record<projects>;
DEFINE FIELD name ON projects TYPE string;
DEFINE FIELD icon ON projects TYPE option<string>;
DEFINE FIELD desc ON projects TYPE string;
DEFINE FIELD owner ON projects TYPE record<users>;
DEFINE FIELD status ON projects TYPE string;
DEFINE FIELD created_at ON projects TYPE datetime VALUE time::now() READONLY;
DEFINE FIELD updated_at ON projects TYPE datetime VALUE time::now() READONLY;
DEFINE FIELD deleted_at ON projects TYPE option<datetime> DEFAULT null;