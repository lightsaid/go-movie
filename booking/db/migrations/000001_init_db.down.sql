
alter table if exists "booking" drop constraint if exists "booking_show_id_fkey";
alter table if exists "booking" drop constraint if exists "booking_user_id_fkey";
alter table if exists "cinema" drop constraint if exists "cinema_city_id_fkey";
alter table if exists "cinema_hall" drop constraint if exists "cinema_hall_cinema_id_fkey";
alter table if exists "cinema_seat" drop constraint if exists "cinema_seat_cinema_hall_id_fkey";
alter table if exists "oauths" drop constraint if exists "oauths_user_id_fkey";
alter table if exists "payment" drop constraint if exists "payment_booking_id_fkey";
alter table if exists "show" drop constraint if exists "show_cinema_hall_id_fkey";
alter table if exists "show" drop constraint if exists "show_movie_id_fkey";
alter table if exists "show_seat" drop constraint if exists "show_seat_booking_fkey";
alter table if exists "show_seat" drop constraint if exists "show_seat_cinema_seat_id_fkey";
alter table if exists "show_seat" drop constraint if exists "show_seat_show_id_fkey";



drop table if exists "user";
drop table if exists "oauths";
drop table if exists "movie";
drop table if exists "city";
drop table if exists "cinema";
drop table if exists "cinema_hall";
drop table if exists "cinema_seat";
drop table if exists "show";
drop table if exists "show_seat";
drop table if exists "booking";
drop table if exists "payment";