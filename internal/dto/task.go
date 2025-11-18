package dto

/*
================================================================================
                                DTO LAYER (READ ME)
================================================================================

WHAT IS A DTO?
---------------
DTO stands for "Data Transfer Object".
A DTO defines the shape of data that flows *into* and *out of* your API.

It is NOT the same as the database model.

Think of DTOs as "API-facing structs",
and Models as "Database-facing structs".

WHY NOT JUST USE MODELS?
------------------------
Your `models/Task` struct represents how data is stored in the database:

    type Task struct {
        Id        string
        Title     string
        Done      bool
        CreatedAt time.Time
    }

But this is NOT what clients send to the server.

For example:
- When creating a task, the client only sends: { "title": "Buy milk" }
  → They should NOT provide: id, done, or created_at.

- When updating a task, the client might only send ONE field:
      { "done": true }
  → Optional fields require pointers (e.g. *string, *bool)

- When responding to the client, you want to control exactly what goes out.
  → Maybe exclude internal fields in the future.

DTOs let you cleanly separate:
- database structure
- request payloads
- response payloads

BENEFITS OF DTOs
----------------
1. Prevents exposing internal database structure to API consumers.
2. Allows flexible request formats (e.g., partial updates with PATCH).
3. Makes validation easier (e.g., title must not be empty).
4. Keeps API stable even if database fields change.
5. Makes your code expressive: "this struct is for requests", "this for responses".

FILES HERE
-----------
CreateTaskRequest
    → Used when client POSTs a new task.

UpdateTaskRequest
    → Used for PATCH /task/{id}.
    → Uses pointers so fields are optional.

TaskResponse
    → What the API sends back after creating, reading, or updating a task.

HOW HANDLERS USE DTOs
----------------------
Example (Create Task):

    var req dto.CreateTaskRequest
    json.NewDecoder(r.Body).Decode(&req)

Example (Returning Response):

    res := dto.TaskResponse{
        Id: task.Id,
        Title: task.Title,
        Done: task.Done,
        CreatedAt: task.CreatedAt,
    }
    RespondWithJSON(w, 200, res)

RELATIONSHIP:
-------------
Client JSON ↔ DTO ↔ Model ↔ DB

DTOs keep each layer cleanly separated.

================================================================================
*/
