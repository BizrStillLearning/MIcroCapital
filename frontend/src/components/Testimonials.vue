<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { Star, X } from '@lucide/vue'

const isVisible = ref(true)
const currentIndex = ref(0)
let intervalId = null

const testimonials = [
  {
    rating: 5,
    text: "Umoja helped me pool funds with my neighbors. Within three months, I was able to buy a new sewing machine for my tailoring business.",
    image: "photo-1531123897727-8f129e1688ce",
    name: "Amina K.",
    role: "Small Business Owner"
  },
  {
    rating: 5,
    text: "The micro-loan process is incredibly fast. I didn't need to visit a bank or fill out endless paperwork. Everything was done on my basic phone.",
    image: "photo-1506277886164-e25aa3f4ef7f",
    name: "David O.",
    role: "Local Farmer"
  },
  {
    rating: 4,
    text: "Creating a savings circle for our community clinic project was so easy. Everyone can see exactly how much we've raised in total transparency.",
    image: "photo-1589156280159-27698a70f29e",
    name: "Sarah M.",
    role: "Community Organizer"
  }
]

const currentTestimonial = computed(() => testimonials[currentIndex.value])

onMounted(() => {
  intervalId = setInterval(() => {
    isVisible.value = false
    setTimeout(() => {
      currentIndex.value = (currentIndex.value + 1) % testimonials.length
      isVisible.value = true
    }, 500)
  }, 5000)
})

onUnmounted(() => {
  if (intervalId) clearInterval(intervalId)
})
</script>

<template>
  <Transition
      enter-active-class="transition duration-500 ease-out"
      enter-from-class="transform translate-y-10 opacity-0"
      enter-to-class="transform translate-y-0 opacity-100"
      leave-active-class="transition duration-300 ease-in"
      leave-from-class="transform translate-y-0 opacity-100"
      leave-to-class="transform translate-y-10 opacity-0"
  >
    <div
        v-if="isVisible"
        class="fixed z-30 bottom-4 left-4 right-4 sm:right-auto sm:bottom-6 sm:left-6 w-auto sm:w-[320px]"
    >
      <div class="bg-card border border-border rounded-2xl p-4 shadow-2xl shadow-black/10 relative group">

        <button
            @click="isVisible = false"
            class="absolute top-3 right-3 text-muted-foreground hover:text-foreground bg-muted/50 hover:bg-muted p-1 rounded-full transition-colors opacity-0 group-hover:opacity-100"
        >
          <X :size="14" />
        </button>

        <div class="flex items-start gap-3">
          <img
              :src="`https://images.unsplash.com/${currentTestimonial.image}?w=50&h=50&fit=crop&auto=format`"
              :alt="currentTestimonial.name"
              class="w-10 h-10 rounded-full object-cover border border-primary/20 flex-shrink-0"
          />

          <div class="flex-1 min-w-0">
            <div class="flex items-center justify-between mb-1 pr-6">
              <p class="text-sm font-bold text-foreground truncate" style="font-family: 'Fraunces', serif;">
                {{ currentTestimonial.name }}
              </p>
              <div class="flex items-center flex-shrink-0">
                <Star v-for="n in currentTestimonial.rating" :key="n" :size="10" class="fill-accent text-accent" />
              </div>
            </div>

            <p class="text-xs text-muted-foreground mb-2" style="font-family: 'DM Mono', monospace;">
              {{ currentTestimonial.role }}
            </p>

            <p class="text-xs text-foreground leading-relaxed italic line-clamp-3" style="font-family: 'DM Sans', sans-serif;">
              "{{ currentTestimonial.text }}"
            </p>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>