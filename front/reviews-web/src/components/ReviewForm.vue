<template>
  <form @submit.prevent="handleSubmit">
    <label>Имя<br>
      <input
        :value="modelValue.name"
        @input="updateField('name', $event.target.value)"
        :readonly="isViewMode"
        :disabled="isViewMode"
        required
      />
    </label><br>

    <label>Дата<br>
      <input
        type="date"
        :value="modelValue.date"
        @input="updateField('date', $event.target.value)"
        :readonly="isViewMode"
        :disabled="isViewMode"
        required
      />
    </label><br>

    <label>Телефон<br>
      <input
        :value="modelValue.phone"
        @input="handlePhoneInput"
        placeholder="89001234567"
        :class="{ 'invalid-phone': phoneError }"
        :readonly="isViewMode"
        :disabled="isViewMode"
        required
      />
    </label><br>

    <label>Email<br>
      <input
        :value="modelValue.email"
        @input="updateField('email', $event.target.value)"
        :readonly="isViewMode"
        :disabled="isViewMode"
        type="email"
        required
      />
    </label><br>

    <label>Технологии:</label><br>
    <div v-for="tech in ['HTML', 'CSS', 'JS', 'С++', 'С#', 'Java']" :key="tech">
      <label>
        <input
          type="checkbox"
          :checked="modelValue.technologies?.includes(tech)"
          @change="toggleTech(tech)"
          :disabled="isViewMode"
        /> {{ tech }}
      </label>
    </div>

    <label>Оценка:</label>
    <StarRating
      :rating="Number(modelValue.rating)"
      @update:rating="updateField('rating', $event)"
      :max-rating="9"
      :increment="1"
      :read-only="isViewMode"
      :star-size="20"
    /><br>

    <label>Комментарий<br>
      <textarea
        :value="modelValue.comment"
        @input="updateField('comment', $event.target.value)"
        :readonly="isViewMode"
        :disabled="isViewMode"
        maxlength="200"
        required
      ></textarea>
    </label>
    <div v-if="!isViewMode">
      {{ 200 - (modelValue.comment?.length || 0) }} символов
    </div><br>

    <button type="button" v-if="isViewMode" @click="$emit('clear')">
      Новый отзыв
    </button>
    <button type="submit" v-else>
      Добавить
    </button>
  </form>
</template>

<script setup>
import { ref } from 'vue'
import StarRating from 'vue-star-rating'

const props = defineProps({
  modelValue: { type: Object, required: true },
  isViewMode: { type: Boolean, default: false }
})

const emit = defineEmits(['update:modelValue', 'submit', 'clear'])

const phoneError = ref(false)

const updateField = (key, value) => {
  if (props.isViewMode) return
  emit('update:modelValue', { ...props.modelValue, [key]: value })
}

const toggleTech = (tech) => {
  if (props.isViewMode) return
  const techs = Array.isArray(props.modelValue.technologies) ? [...props.modelValue.technologies] : []
  if (techs.includes(tech)) {
    updateField('technologies', techs.filter(t => t !== tech))
  } else {
    updateField('technologies', [...techs, tech])
  }
}

const handlePhoneInput = (event) => {
  if (props.isViewMode) return
  const phone = event.target.value
  updateField('phone', phone)
  const valid = /^[78]\d{10}$/.test(phone.trim())
  phoneError.value = !valid
}

const handleSubmit = () => {
  if (props.isViewMode) return
  emit('submit', props.modelValue)
}
</script>

<style scoped>

.invalid-phone { border: 1px solid red; }
</style>