<template>
  <div>
    <h2>Отзывы</h2>
    <ReviewList :reviews="reviews" :selectedReview="selectedReview" @select="selectReview" />

    <h2>{{ selectedReview ? 'Просмотр отзыва' : 'Добавить новый отзыв' }}</h2>
    <form @submit.prevent="addReview">
      <label>Имя<br>
        <input v-model="form.name" :readonly="!!selectedReview" required>
      </label><br>
      <label>Дата<br>
        <input type="date" v-model="form.date" :readonly="!!selectedReview" required>
      </label><br>
      <label>Телефон<br>
        <input
        v-model="form.phone"
        :readonly="!!selectedReview"
        @input="isValidPhone"
        @blur="isValidPhone"
        placeholder="89001234567"
        :class="{ 'invalid-phone': phoneError }"
        required
      />
      </label><br>
      <label>Email<br>
        <input v-model="form.email" :readonly="!!selectedReview" required>
      </label><br>

      <label>Технологии:</label>
      <label><input type="checkbox" v-model="form.technologies" value="HTML" :disabled="!!selectedReview"> HTML</label>
      <label><input type="checkbox" v-model="form.technologies" value="CSS" :disabled="!!selectedReview"> CSS</label>
      <label><input type="checkbox" v-model="form.technologies" value="JS" :disabled="!!selectedReview"> JS</label>
      <label><input type="checkbox" v-model="form.technologies" value="С++" :disabled="!!selectedReview"> С++</label>
      <label><input type="checkbox" v-model="form.technologies" value="С#" :disabled="!!selectedReview"> С#</label>
      <label><input type="checkbox" v-model="form.technologies" value="Java" :disabled="!!selectedReview"> Java</label><br>

      <label>Оценка:</label>
      <star-rating v-model:rating="form.rating" :max-rating="9" :increment="1" :read-only="!!selectedReview" :star-size="20" />

      <label>Комментарий<br>
        <textarea v-model="form.comment" :readonly="!!selectedReview" maxlength="200" required></textarea>
      </label>
      <div>{{ 200 - form.comment.length }} символов</div><br>

      <button type="button" v-if="selectedReview" @click="clearForm">Новый отзыв</button>
      <button type="submit" v-else>Добавить</button>
    </form>

    <p><strong>Средний рейтинг:</strong> {{ averageRating.toFixed(2) }}</p>
    <ReviewDetail v-if="selectedReview" :review="selectedReview" />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import ReviewList from './components/ReviewList.vue'
import StarRating from 'vue-star-rating'

const reviews = ref([])
const selectedReview = ref(null)

const form = ref({
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



const clearForm = () => {
  selectedReview.value = null
  form.value = { name: '', date: '', phone: '', email: '', technologies: [], rating: 0, comment: '' }
}


const phoneError = ref(false)

const isValidPhone = () => {

  const phone = form.value.phone.trim()

  if (phone === '') {
    phoneError.value = false
    return false
  }

  const valid = /^[78]\d{10}$/.test(phone)
  phoneError.value = !valid
  return valid
}
const addReview = () => {
  if (!isValidPhone()) {
    return
  }

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
        throw new Error()
      }
    })
    .catch(() => alert('Не удалось отправить отзыв'))
}


const selectReview = (review) => {
  selectedReview.value = review
  form.value = { ...review }
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
    color: red;
  }
</style>