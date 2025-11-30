<template>
  <table>
    <thead>
      <tr>
        <th>Имя</th>
        <th>Дата рождения</th>
        <th>Технологии</th>
        <th>Оценка</th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="(review, index) in reviews"
        :key="index"
        @click="emit('select', review)"
        style="cursor: pointer"
      >
        <td>{{ review.name }}</td>
        <td>{{ review.date }}</td>
        <td>{{ review.technologies.join(', ') }}</td>
        <td>{{ review.rating }}</td>
      </tr>
    </tbody>
  </table>
</template>

<script setup>
const props = defineProps(['reviews'])
const emit = defineEmits(['select'])

</script>