
# Introduction

This project is an accounting journal application where users can insert data to the database and show various graphs showing the amount of expenses and profit of the company.

Features:

- Different user permissions; employees, team leader, and admin.
- Users are grouped by sections/teams.
- Admin/Team leaders can add and assign employees.
- Insert and retrieve data.
- Upload and retrieve receipt images.
- A dashboard of expenses and income, which can be customized.
- Creates a printable monthly report following GAAP.

The accounting journal within the database are divided into two tables: expense and incom.But they are generally formated as the following:

- `type`: Type of transactions using ID's for example; 101 can mean sales or 102 can mean payments.
- `section`: This field will automatically be filled in based on the user section/team.
- `total`: The total amount of the transaction.
- `date`: The date of the transaction.
- `description`: Additional description for the transaction.

You can also add receipt images if needed, and it will be saved binded to the transaction.
