<template>
  <div class="container">
    <h2>Добавить новый отзыв</h2>
    <form @submit.prevent="addReview" class="add-review-form">
      
      <div class="form-column">
        <div class="form-group">
          <label>Имя</label>
          <input v-model="newReview.name" required />
        </div>

        <div class="form-group">
          <label>Дата рождения</label>
          <input type="date" v-model="newReview.date" required />
        </div>
      </div>

      <div class="form-column">
        <div class="form-group">
          <label>Телефон</label>
          <input v-model="newReview.phone" required placeholder="8XXXXXXXXXX" />
        </div>

        <div class="form-group">
          <label>Email</label>
          <input v-model="newReview.email" required />
        </div>
      </div>

      <div class="form-group">
        <label>Технологии:</label>
        <br />
        <label><input type="checkbox" v-model="newReview.technologies" value="HTML" /> HTML</label>
        <label><input type="checkbox" v-model="newReview.technologies" value="CSS" /> CSS</label>
        <label><input type="checkbox" v-model="newReview.technologies" value="JS" /> JS</label>
        <label><input type="checkbox" v-model="newReview.technologies" value="С++" /> C++</label>
        <label><input type="checkbox" v-model="newReview.technologies" value="С#" /> C#</label>
        <label><input type="checkbox" v-model="newReview.technologies" value="Java" /> Java</label>
      </div>

      <div class="form-group">
        <label>Моя оценка:</label>
        <div>
          <star-rating
            v-model:rating="newReview.rating"
            :max-rating="9"
            :increment="1"
            :star-size="24"
            :show-rating="true"
            :fixed-points="0"
            inactive-color="#ddd"
            active-color="#ffd700"
          />
        </div>
      </div>

      <div class="form-group">
        <label>Комментарий</label><br>
        <textarea
          v-model="newReview.comment"
          rows="4"
          maxlength="200"
          required
        ></textarea>
        <p class="char-counter">{{ 200 - newReview.comment.length }} символов осталось</p>
      </div>

      <button type="submit">Добавить отзыв</button>
      <hr />
    </form>


    <p><strong>Средняя оценка сайта:</strong> {{ averageRating.toFixed(2) }}</p>

    <ReviewList :reviews="reviews" @select="toggleReview" />

    <ReviewDetail v-if="selectedReview" :review="selectedReview" />

  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import ReviewList from './components/ReviewList.vue'
import ReviewDetail from './components/ReviewFull.vue'
import StarRating from 'vue-star-rating'


const reviews = ref([])
const selectedReview = ref(null)
const newReview = ref({
  name: "",
  date: "",
  phone: "",
  email: "",
  technologies: [],
  rating: 0,
  comment: ""
})


const averageRating = computed(() => {
  if (reviews.value.length === 0) return 0
  return (reviews.value.reduce((sum, r) => sum + r.rating, 0) / reviews.value.length)
})


const loadReviews = async () => {
  try {
    const res = await fetch('http://localhost:8082/send')
    if (!res.ok) throw new Error('Не удалось загрузить отзывы')
    reviews.value = await res.json()
  } catch (err) {
    console.error('Ошибка загрузки:', err)
    alert('Не удалось загрузить отзывы')
  }
}


const addReview = async () => {
  try {
    const res = await fetch('http://localhost:8082/save', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newReview.value)
    })

    if (!res.ok) {
      const err = await res.text().catch(() => 'Неизвестная ошибка')
      throw new Error(err)
    }

  
    await loadReviews()
    newReview.value = {
      name: "",
      date: "",
      phone: "",
      email: "",
      technologies: [],
      rating: 0,
      comment: ""
    }
    selectedReview.value = null
  } catch (err) {
    console.error('Ошибка отправки:', err)
    alert('Не удалось отправить отзыв: ' + err.message)
  }
}

const toggleReview = (review) => {
  selectedReview.value = selectedReview.value === review ? null : review
}


onMounted(() => {
  loadReviews()
})
</script>

<style>



.form-column{
  display: inline-block;
}
label {
  display: inline-block;
  width: 180px;
  vertical-align: top;
}
button {
  margin-top: 10px;
  padding: 8px 16px;
  background: #ce4a1a;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
textarea {
  resize: none;
  width: 400px;
  height: 200px;
  border-radius: 3px;
}
th, td {
  border: 1px solid #ccc;
  padding: 10px;
  text-align: left;
}
</style>