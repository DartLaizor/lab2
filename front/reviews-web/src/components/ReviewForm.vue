<template>
  <form @submit.prevent="handleSubmit">
    <label>Имя<br>
      <input 
        :value="modelValue.name" 
        @input="updateField('name', $event.target.value)" 
        :readonly="!!modelValue.id" 
        required
      >
    </label><br>

    <label>Дата<br>
      <input 
        type="date" 
        :value="modelValue.date" 
        @input="updateField('date', $event.target.value)" 
        :readonly="!!modelValue.id" 
        required
      >
    </label><br>

    <label>Телефон<br>
      <input
        :value="modelValue.phone"
        @input="handlePhoneInput"
        @blur="validatePhone(modelValue.phone)"
        placeholder="89001234567"
        :class="{ 'invalid-phone': phoneError }"
        :readonly="!!modelValue.id"
        required
      />
    </label><br>

    <label>Email<br>
      <input 
        :value="modelValue.email" 
        @input="updateField('email', $event.target.value)" 
        :readonly="!!modelValue.id" 
        required
      >
    </label><br>

    <label>Технологии:</label>
    <label>
      <input 
        type="checkbox" 
        :checked="modelValue.technologies.includes('HTML')" 
        @change="toggleTech('HTML')"
        :disabled="!!modelValue.id"
      > HTML
    </label>
    <label>
      <input 
        type="checkbox" 
        :checked="modelValue.technologies.includes('CSS')" 
        @change="toggleTech('CSS')"
        :disabled="!!modelValue.id"
      > CSS
    </label>
    <label>
      <input 
        type="checkbox" 
        :checked="modelValue.technologies.includes('JS')" 
        @change="toggleTech('JS')"
        :disabled="!!modelValue.id"
      > JS
    </label>
    <label>
      <input 
        type="checkbox" 
        :checked="modelValue.technologies.includes('С++')" 
        @change="toggleTech('С++')"
        :disabled="!!modelValue.id"
      > С++
    </label>
    <label>
      <input 
        type="checkbox" 
        :checked="modelValue.technologies.includes('С#')" 
        @change="toggleTech('С#')"
        :disabled="!!modelValue.id"
      > С#
    </label>
    <label>
      <input 
        type="checkbox" 
        :checked="modelValue.technologies.includes('Java')" 
        @change="toggleTech('Java')"
        :disabled="!!modelValue.id"
      > Java
    </label><br>

    <label>Оценка:</label>
    <StarRating 
      :rating="modelValue.rating" 
      @update:rating="updateField('rating', $event)" 
      :max-rating="9" 
      :increment="1" 
      :read-only="!!modelValue.id" 
      :star-size="20" 
    />

    <label>Комментарий<br>
      <textarea 
        :value="modelValue.comment" 
        @input="updateField('comment', $event.target.value)" 
        :readonly="!!modelValue.id" 
        maxlength="200" 
        required
      ></textarea>
    </label>
    <div>{{ 200 - modelValue.comment.length }} символов</div><br>

    <button 
      type="button" 
      v-if="modelValue.id" 
      @click="handleClear"
    >
      Новый отзыв
    </button>
    <button type="submit" v-else>Добавить</button>
  </form>
</template>

<script setup>
import { ref } from 'vue'
import StarRating from 'vue-star-rating'

const props = defineProps({
  modelValue: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['update:modelValue', 'submit', 'clear'])

const phoneError = ref(false)

const updateField = (key, value) => {
  emit('update:modelValue', { ...props.modelValue, [key]: value })
}

const toggleTech = (tech) => {
  const techs = props.modelValue.technologies || []
  if (techs.includes(tech)) {
    updateField('technologies', techs.filter(t => t !== tech))
  } else {
    updateField('technologies', [...techs, tech])
  }
}

const validatePhone = (phone) => {
  if (phone === '') {
    phoneError.value = false
    return true
  }
  const valid = /^[78]\d{10}$/.test(phone.trim())
  phoneError.value = !valid
  return valid
}

const handlePhoneInput = (event) => {
  const phone = event.target.value
  updateField('phone', phone)
  validatePhone(phone)
}

const handleSubmit = () => {
  if (!validatePhone(props.modelValue.phone)) return
  emit('submit')
}

const handleClear = () => {
  phoneError.value = false
  emit('clear')
}
</script>

<style scoped>
.invalid-phone {
  border: 2px solid red;
}
</style>