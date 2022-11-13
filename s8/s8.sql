with redSales as (
select 
    a.sales_id
from 
    Orders a
    join Company b on a.com_id = b.com_id and b.name = 'RED'
)

select name from SalesPerson
where sales_id not in (select * from redSales);
