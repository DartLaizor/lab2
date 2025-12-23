<template>
  <div>
    <h2>Отзывы</h2>
    <ReviewList 
      :reviews="reviews" 
      :selectedReview="selectedReview" 
      @select="selectReview" 
    />

    <h2>{{ selectedReview ? 'Просмотр отзыва' : 'Добавить новый отзыв' }}</h2>
    
    <ReviewForm
      v-model="form"
      :is-view-mode="!!selectedReview"
      @submit="addReview"
      @clear="clearForm"
    />

    <p><strong>Средний рейтинг:</strong> {{ averageRating.toFixed(2) }}</p>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import ReviewForm from './components/ReviewForm.vue'
import ReviewList from './components/ReviewList.vue'

const reviews = ref([])
const selectedReview = ref(null)

const getInitialForm = () => ({
  name: '',
  date: '',
  phone: '',
  email: '',
  technologies: [],
  rating: 0,
  comment: ''
})

const form = ref(getInitialForm())

const averageRating = computed(() => {
  const total = reviews.value.reduce((s, r) => s + (Number(r.rating) || 0), 0)
  return reviews.value.length ? total / reviews.value.length : 0
})

const addReview = (reviewData) => {
  fetch('http://localhost:8082/save', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(reviewData)
  })
    .then(res => {
      if (res.ok) {
        loadReviews()
        clearForm()
      } else {
        alert('Ошибка сохранения')
      }
    })
    .catch(() => alert('Не удалось отправить отзыв'))
}

const selectReview = (review) => {
  selectedReview.value = review
  form.value = { ...review }
}

const clearForm = () => {
  selectedReview.value = null
  form.value = getInitialForm()
}

const loadReviews = () => {
  fetch('http://localhost:8082/send')
    .then(res => res.json())
    .then(data => reviews.value = Array.isArray(data) ? data : [])
    .catch(() => alert('Ошибка загрузки'))
}

onMounted(loadReviews)
</script>