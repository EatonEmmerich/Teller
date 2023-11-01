create table public.invoice
(
    id integer not null
);

create sequence invoice_id_seq owned by invoice.id;