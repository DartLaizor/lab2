<template>
  <table>
    <thead>
      <tr>
        <th>Имя</th>
        <th>Дата</th>
        <th>Технологии</th>
        <th>Оценка</th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="(r, i) in reviews"
        :key="i"
        @click="emit('select', r)"
        :class="{ selected: r === selectedReview }"
        style="cursor: pointer"
      >
        <td>{{ r.name }}</td>
        <td>{{ r.date }}</td>
        <td>{{ r.technologies?.join(', ') }}</td>
        <td>{{ r.rating }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script setup>
defineProps(['reviews', 'selectedReview'])
const emit = defineEmits(['select'])
</script>

<style scoped>
table { width: 100%; border-collapse: collapse; margin-bottom: 1rem; }
th, td { border: 1px solid #ccc; padding: 8px; text-align: left; }
.selected { background-color: #e0e0e0; }
</style>