import React, { useState, useEffect } from 'react';
import { 
  Typography, 
  Box, 
  Paper, 
  List, 
  ListItem, 
  ListItemText, 
  ListItemSecondaryAction, 
  IconButton, 
  TextField, 
  Button, 
  Checkbox, 
  Divider, 
  CircularProgress,
  Alert,
  Grid
} from '@mui/material';
import { Delete as DeleteIcon, Edit as EditIcon, Save as SaveIcon, Cancel as CancelIcon } from '@mui/icons-material';
import axios from 'axios';

const TodoList = () => {
  const [todos, setTodos] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [newTodo, setNewTodo] = useState({ title: '', description: '' });
  const [editingTodo, setEditingTodo] = useState(null);
  const [editFormData, setEditFormData] = useState({ title: '', description: '', completed: false });

  // Fetch todos on component mount
  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    try {
      setLoading(true);
      setError('');
      const response = await axios.get('/api/todos');
      setTodos(response.data);
    } catch (err) {
      console.error('Error fetching todos:', err);
      setError('Failed to load todos. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const handleNewTodoChange = (e) => {
    const { name, value } = e.target;
    setNewTodo(prev => ({ ...prev, [name]: value }));
  };

  const handleEditFormChange = (e) => {
    const { name, value, type, checked } = e.target;
    setEditFormData(prev => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : value
    }));
  };

  const handleAddTodo = async (e) => {
    e.preventDefault();
    if (!newTodo.title.trim()) {
      setError('Title is required');
      return;
    }

    try {
      setLoading(true);
      setError('');
      await axios.post('/api/todos', newTodo);
      setNewTodo({ title: '', description: '' });
      await fetchTodos();
    } catch (err) {
      console.error('Error adding todo:', err);
      setError('Failed to add todo. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const handleUpdateTodo = async (e) => {
    e.preventDefault();
    if (!editFormData.title.trim()) {
      setError('Title is required');
      return;
    }

    try {
      setLoading(true);
      setError('');
      await axios.put(`/api/todos/${editingTodo.id}`, editFormData);
      setEditingTodo(null);
      await fetchTodos();
    } catch (err) {
      console.error('Error updating todo:', err);
      setError('Failed to update todo. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const handleDeleteTodo = async (id) => {
    try {
      setLoading(true);
      setError('');
      await axios.delete(`/api/todos/${id}`);
      await fetchTodos();
    } catch (err) {
      console.error('Error deleting todo:', err);
      setError('Failed to delete todo. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const handleToggleComplete = async (todo) => {
    try {
      setLoading(true);
      setError('');
      await axios.put(`/api/todos/${todo.id}`, {
        completed: !todo.completed
      });
      await fetchTodos();
    } catch (err) {
      console.error('Error updating todo:', err);
      setError('Failed to update todo. Please try again.');
    } finally {
      setLoading(false);
    }
  };

  const startEditing = (todo) => {
    setEditingTodo(todo);
    setEditFormData({
      title: todo.title,
      description: todo.description || '',
      completed: todo.completed
    });
  };

  const cancelEditing = () => {
    setEditingTodo(null);
    setEditFormData({ title: '', description: '', completed: false });
  };

  // Filter todos into complete and incomplete
  const completedTodos = todos.filter(todo => todo.completed);
  const incompleteTodos = todos.filter(todo => !todo.completed);

  // Render todo item (used for both complete and incomplete lists)
  const renderTodoItem = (todo) => (
    <React.Fragment key={todo.id}>
      {editingTodo && editingTodo.id === todo.id ? (
        <ListItem>
          <Box component="form" onSubmit={handleUpdateTodo} sx={{ width: '100%' }}>
            <TextField
              margin="dense"
              required
              fullWidth
              name="title"
              label="Title"
              value={editFormData.title}
              onChange={handleEditFormChange}
            />
            <TextField
              margin="dense"
              fullWidth
              name="description"
              label="Description"
              multiline
              rows={2}
              value={editFormData.description}
              onChange={handleEditFormChange}
            />
            <Box sx={{ display: 'flex', alignItems: 'center', mt: 1 }}>
              <Checkbox
                checked={editFormData.completed}
                onChange={handleEditFormChange}
                name="completed"
              />
              <Typography variant="body2">Completed</Typography>
            </Box>
            <Box sx={{ mt: 1, display: 'flex', gap: 1 }}>
              <Button
                type="submit"
                variant="contained"
                size="small"
                startIcon={<SaveIcon />}
                disabled={loading}
              >
                Save
              </Button>
              <Button
                variant="outlined"
                size="small"
                startIcon={<CancelIcon />}
                onClick={cancelEditing}
                disabled={loading}
              >
                Cancel
              </Button>
            </Box>
          </Box>
        </ListItem>
      ) : (
        <ListItem>
          <Checkbox
            edge="start"
            checked={todo.completed}
            onChange={() => handleToggleComplete(todo)}
            disabled={loading}
          />
          <ListItemText
            primary={
              <Typography
                variant="body1"
                style={{
                  textDecoration: todo.completed ? 'line-through' : 'none',
                  color: todo.completed ? 'text.secondary' : 'text.primary'
                }}
              >
                {todo.title}
              </Typography>
            }
            secondary={todo.description}
          />
          <ListItemSecondaryAction>
            <IconButton
              edge="end"
              aria-label="edit"
              onClick={() => startEditing(todo)}
              disabled={loading}
            >
              <EditIcon />
            </IconButton>
            <IconButton
              edge="end"
              aria-label="delete"
              onClick={() => handleDeleteTodo(todo.id)}
              disabled={loading}
            >
              <DeleteIcon />
            </IconButton>
          </ListItemSecondaryAction>
        </ListItem>
      )}
      <Divider />
    </React.Fragment>
  );

  return (
    <Box>
      <Typography variant="h4" component="h1" gutterBottom>
        My Todo List
      </Typography>

      {error && (
        <Alert severity="error" sx={{ mb: 2 }}>
          {error}
        </Alert>
      )}

      {/* Add new todo form */}
      <Paper elevation={3} sx={{ p: 3, mb: 4 }}>
        <Typography variant="h6" gutterBottom>
          Add New Todo
        </Typography>
        <Box component="form" onSubmit={handleAddTodo}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="title"
            label="Title"
            name="title"
            value={newTodo.title}
            onChange={handleNewTodoChange}
          />
          <TextField
            margin="normal"
            fullWidth
            id="description"
            label="Description"
            name="description"
            multiline
            rows={3}
            value={newTodo.description}
            onChange={handleNewTodoChange}
          />
          <Button
            type="submit"
            variant="contained"
            sx={{ mt: 2 }}
            disabled={loading}
          >
            {loading ? <CircularProgress size={24} /> : 'Add Todo'}
          </Button>
        </Box>
      </Paper>

      {/* Todo lists in grid layout */}
      <Grid container spacing={3}>
        {/* Incomplete Todos Card */}
        <Grid item xs={12} md={6}>
          <Paper elevation={3} sx={{ p: 0, height: '100%' }}>
            <Box sx={{ p: 2, bgcolor: 'primary.main', color: 'white' }}>
              <Typography variant="h6">Incomplete Todos</Typography>
            </Box>
            <List>
              {loading && todos.length === 0 ? (
                <Box sx={{ display: 'flex', justifyContent: 'center', p: 3 }}>
                  <CircularProgress />
                </Box>
              ) : incompleteTodos.length === 0 ? (
                <ListItem>
                  <ListItemText primary="No incomplete todos. Great job!" />
                </ListItem>
              ) : (
                incompleteTodos.map(renderTodoItem)
              )}
            </List>
          </Paper>
        </Grid>

        {/* Completed Todos Card */}
        <Grid item xs={12} md={6}>
          <Paper elevation={3} sx={{ p: 0, height: '100%' }}>
            <Box sx={{ p: 2, bgcolor: 'success.main', color: 'white' }}>
              <Typography variant="h6">Completed Todos</Typography>
            </Box>
            <List>
              {loading && todos.length === 0 ? (
                <Box sx={{ display: 'flex', justifyContent: 'center', p: 3 }}>
                  <CircularProgress />
                </Box>
              ) : completedTodos.length === 0 ? (
                <ListItem>
                  <ListItemText primary="No completed todos yet. Keep going!" />
                </ListItem>
              ) : (
                completedTodos.map(renderTodoItem)
              )}
            </List>
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
};

export default TodoList;