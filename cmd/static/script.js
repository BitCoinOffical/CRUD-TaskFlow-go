const API_URL = 'http://localhost:8080'; // Замени на свой URL

// DOM-элементы
const taskForm = document.getElementById('taskForm');
const taskList = document.getElementById('taskList');
const searchInput = document.getElementById('search');
const filterPriority = document.getElementById('filterPriority');

// Загрузка задач при старте
document.addEventListener('DOMContentLoaded', fetchTasks);

// Добавление/обновление задачи
taskForm.addEventListener('submit', async (e) => {
  e.preventDefault();
  
  const task = {
    title: document.getElementById('title').value,
    description: document.getElementById('description').value,
    priority: document.getElementById('priority').value,
    status: document.getElementById('status').value
  };
  
  try {
    const response = await fetch(`${API_URL}/tasks`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(task)
    });
    
    if (response.ok) {
      fetchTasks();
      taskForm.reset();
    }
  } catch (error) {
    console.error('Ошибка:', error);
  }
});

// Получение списка задач
async function fetchTasks() {
  try {
    const response = await fetch(`${API_URL}/tasks`);
    const tasks = await response.json();
    renderTasks(tasks);
  } catch (error) {
    console.error('Ошибка загрузки задач:', error);
  }
}

// Фильтрация и поиск
searchInput.addEventListener('input', filterTasks);
filterPriority.addEventListener('change', filterTasks);

function filterTasks() {
  const searchTerm = searchInput.value.toLowerCase();
  const priority = filterPriority.value;
  
  fetch(`${API_URL}/tasks`)
    .then(res => res.json())
    .then(tasks => {
      const filtered = tasks.filter(task => {
        const matchesSearch = task.title.toLowerCase().includes(searchTerm);
        const matchesPriority = priority ? task.priority === priority : true;
        return matchesSearch && matchesPriority;
      });
      renderTasks(filtered);
    });
}

// Отображение задач
function renderTasks(tasks) {
  taskList.innerHTML = tasks.map(task => `
    <div class="task ${task.status === 'completed' ? 'completed' : ''}">
      <div>
        <h3>${task.title}</h3>
        <p>${task.description}</p>
        <span class="priority-${task.priority}">
          Приоритет: ${task.priority} | Статус: ${task.status}
        </span>
      </div>
      <div class="task-actions">
        <button onclick="updateTaskStatus('${task.id}', 'completed')">✓</button>
        <button onclick="deleteTask('${task.id}')" class="delete">✕</button>
      </div>
    </div>
  `).join('');
}

// Удаление задачи
async function deleteTask(id) {
  if (confirm('Удалить задачу?')) {
    await fetch(`${API_URL}/tasks/${id}`, { method: 'DELETE' });
    fetchTasks();
  }
}

// Обновление статуса
async function updateTaskStatus(id, status) {
  await fetch(`${API_URL}/tasks/${id}`, {
    method: 'PATCH',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ status })
  });
  fetchTasks();
}