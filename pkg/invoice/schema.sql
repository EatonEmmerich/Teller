create table public.invoice
(
    id integer not null,
    paid bool default false
);

create sequence invoice_id_seq owned by invoice.id;

create table public.invoice_item
(
    invoice_id integer not null,
    amount integer not null,
    description text
);