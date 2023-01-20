Project Checklist

- A Headless CMS dashboard
  - Authentication to enter the dashboard
  - Server rendered UI with direct access to Go functions
  - Section where you can create a db table and define its columns (and respective types)
    - The columns must be among predefined options, such as text (string), integer (int), etc.
    - Each new column created must be specified as a required field or not
    - When a new column is created, it must ship two unmutable fields: ID and date created
    - Every field and the table itself must have CRUD options
  - Section where you can create entries in a predefined schema
    - Allow unlimited CRUD operations
  - Autogenerated GET APIs of the content created
    - The generated API must be structured as websiteURLexample.com/api/tableName
    - Query params must be available to filter results
  