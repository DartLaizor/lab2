<template>
    <div class = "full-review">
        <p><strong>Имя:</strong> {{ review.name }}</p>
        <p><strong>Дата рождения:</strong> {{review.date }}</p>
        <p><strong>Телефон:</strong> {{ review.phone }}</p>
        <p><strong>Email:</strong> {{ review.email }}</p>
        <p><strong>Изучаемые технологии:</strong> {{ review.technologies.join(', ') }}</p>
        <p><strong>Оценка сайта:</strong> {{ review.rating }}/9</p>
        <p><strong>Отзыв:</strong><br />{{ review.comment }}</p>
    </div>
</template>

<script setup>
defineProps(['review'])
</script>