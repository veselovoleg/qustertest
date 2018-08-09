# qustertest
quster

#postgres

CREATE TABLE public.questions (
	id int4 NOT NULL,
	question varchar NOT NULL,
	rightanswer varchar NOT NULL,
	CONSTRAINT questions_pkey PRIMARY KEY (id)
)
