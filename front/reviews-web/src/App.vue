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
      :phone-error="phoneError"
      @validate-phone="validatePhone"
      @submit="addReview"
      @clear="clearForm"
    />

    <p><strong>Средний рейтинг:</strong> {{ averageRating.toFixed(2) }}</p>
    
    <ReviewDetail v-if="selectedReview" :review="selectedReview" />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import ReviewForm from './components/ReviewForm.vue'
import ReviewList from './components/ReviewList.vue'


const reviews = ref([])
const selectedReview = ref(null)
const phoneError = ref(false)

const form = ref({
  id: null,
  name: '',
  date: '',
  phone: '',
  email: '',
  technologies: [],
  rating: 0,
  comment: ''
})

const averageRating = computed(() => {
  const total = reviews.value.reduce((s, r) => s + r.rating, 0)
  return reviews.value.length ? total / reviews.value.length : 0
})

const validatePhone = (phone) => {
  if (phone === '') {
    phoneError.value = false
    return true
  }
  const valid = /^[78]\d{10}$/.test(phone.trim())
  phoneError.value = !valid
  return valid
}

const addReview = () => {
  if (!validatePhone(form.value.phone)) return

  fetch('http://localhost:8082/save', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(form.value)
  })
    .then(res => {
      if (res.ok) {
        loadReviews()
        clearForm()
      } else {
        throw new Error('HTTP error')
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
  form.value = {
    id: null,
    name: '',
    date: '',
    phone: '',
    email: '',
    technologies: [],
    rating: 0,
    comment: ''
  }
  phoneError.value = false
}

const loadReviews = () => {
  fetch('http://localhost:8082/send')
    .then(res => res.json())
    .then(data => reviews.value = Array.isArray(data) ? data : [])
    .catch(() => alert('Ошибка загрузки отзывов'))
}

onMounted(loadReviews)
</script>

<style>
.invalid-phone {
  border: 2px solid red;
}
</style>