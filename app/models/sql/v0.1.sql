create table groups(
groupid serial primary key, 
title varchar );

create table tasks(
taskid serial primary key, 
title varchar, 
groupid int references groups (groupid) on delete cascade);

create table timeframes(
taskid int references task(taskid), 
from timestamp without time zone,
to timestamp without time zone);