<template>
  <div class="container">
    <h2>Отзывы</h2>
    <ReviewList :reviews="reviews" :selectedReview="selectedReview" @select="toggleReview" />

    <h2 style="margin-top: 2rem">{{ isViewMode ? 'Просмотр отзыва' : 'Добавить новый отзыв' }}</h2>
    <form @submit.prevent="addReview" class="add-review-form">
      <div class="form-row">
        <div>
          <label>Имя</label>
          <input v-model="formData.name" :readonly="isViewMode" required />
        </div>
        <div>
          <label>Дата</label>
          <input type="date" v-model="formData.date" :readonly="isViewMode" required />
        </div>
        <div>
          <label>Телефон</label>
          <input
            v-model="formData.phone"
            :readonly="isViewMode"
            @input="validatePhone"
            placeholder="89001234567"
            required
          />
          <div v-if="phoneError" class="error">{{ phoneError }}</div>
        </div>
        <div>
          <label>Email</label>
          <input v-model="formData.email" :readonly="isViewMode" required />
        </div>
      </div>

      <div class="tech-group">
        <label>Технологии:</label>
        <label v-for="tech in techList" :key="tech">
          <input
            type="checkbox"
            v-model="formData.technologies"
            :value="tech"
            :disabled="isViewMode"
          />
          {{ tech }}
        </label>
      </div>

      <div>
        <label>Оценка:</label>
        <star-rating
          v-model:rating="formData.rating"
          :max-rating="9"
          :increment="1"
          :read-only="isViewMode"
          :star-size="24"
          inactive-color="#ddd"
          active-color="#ffd700"
        />
      </div>

      <div>
        <label>Комментарий</label>
        <textarea
          v-model="formData.comment"
          :readonly="isViewMode"
          maxlength="200"
          required
        ></textarea>
        <div class="char-counter">{{ 200 - formData.comment.length }} символов</div>
      </div>

      <button type="button" v-if="isViewMode" @click="resetForm">Новый отзыв</button>
      <button type="submit" v-else>Добавить</button>
    </form>

    <p><strong>Средний рейтинг:</strong> {{ averageRating.toFixed(2) }}</p>
    <ReviewDetail v-if="selectedReview" :review="selectedReview" />
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import ReviewList from './components/ReviewList.vue'
import ReviewDetail from './components/ReviewDetail.vue'
import StarRating from 'vue-star-rating'

const reviews = ref([])
const selectedReview = ref(null)
const newReview = ref({ name: '', date: '', phone: '', email: '', technologies: [], rating: 0, comment: '' })
const phoneError = ref('')

const techList = ['HTML', 'CSS', 'JS', 'С++', 'С#', 'Java']

const isViewMode = computed(() => !!selectedReview.value)
const formData = computed(() => isViewMode.value ? selectedReview.value : newReview.value)
const averageRating = computed(() => reviews.value.reduce((s, r) => s + r.rating, 0) / (reviews.value.length || 1))

const validatePhone = () => {
  const d = newReview.value.phone.replace(/\D/g, '')
  phoneError.value = !d ? 'Телефон обязателен' : (d.length !== 11 || !/^[78]/.test(d)) ? 'Формат: 89001234567' : ''
}

watch(isViewMode, () => phoneError.value = '')

const resetForm = () => {
  selectedReview.value = null
  newReview.value = { name: '', date: '', phone: '', email: '', technologies: [], rating: 0, comment: '' }
}

const loadReviews = async () => {
  try {
    const res = await fetch('http://localhost:8082/send')
    reviews.value = res.ok ? await res.json() : []
  } catch (e) {
    alert('Ошибка загрузки отзывов')
  }
}

const addReview = async () => {
  validatePhone()
  if (phoneError.value) return alert('Исправьте телефон')
  try {
    const res = await fetch('http://localhost:8082/save', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newReview.value)
    })
    if (res.ok) {
      await loadReviews()
      resetForm()
    } else throw new Error()
  } catch {
    alert('Не удалось отправить отзыв')
  }
}

const toggleReview = (review) => {
  selectedReview.value = selectedReview.value === review ? null : review
}

onMounted(loadReviews)
</script>

<style scoped>
.container {
  max-width: 900px;
  margin: 0 auto;
  font-family: sans-serif;
}
.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}
.form-row > div {
  display: flex;
  flex-direction: column;
}
label { font-weight: bold; margin-bottom: 4px; }
input, textarea {
  padding: 6px;
  font-size: 14px;
  border: 1px solid #ccc;
  border-radius: 4px;
}
textarea {
  width: 100%;
  height: 100px;
  resize: vertical;
}
.tech-group {
  margin: 12px 0;
}
.tech-group label {
  margin-right: 12px;
  font-weight: normal;
}
.error { color: red; font-size: 12px; margin-top: 4px; }
.char-counter { font-size: 12px; color: #666; margin-top: 4px; }
button {
  margin-top: 12px;
  padding: 6px 16px;
  background: #ce4a1a;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
</style>