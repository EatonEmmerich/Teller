create table public.payment
(
    id integer not null,
    amount integer not null,
    invoice integer not null
);

create sequence payment_id_seq owned by payment.id;