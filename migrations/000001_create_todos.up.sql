CREATE TABLE todos (
	id SERIAL PRIMARY KEY,
	title VARCHAR(100) NOT NULL,
	status VARCHAR(20) NOT NULL CHECK (status IN ('todo', 'done'))
);