# cs2-panel

***! This is a temporary documentation, meant as a general overview***

### Permission System
#### General
- Every interaction with the server is checked against the users permissions
- Roles can be assigned specific permissions for every relevant action
- A permission can have a parent permission, which toggles all it's children
- Users can have multiple roles

#### Commands
- Only Admins can add new and edit existing commands
- Every command has a category
- Every command has a permission
- Every category has a permission, which acts as the parent permission of the categorys commands

Permission format: ``exec_{command}``, example: ``exec_mp_warmup_end``
