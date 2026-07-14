<script setup>
import { Star } from '@lucide/vue'

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
</script>

<template>
  <section id="community" class="py-20 md:py-28">
    <div class="max-w-6xl mx-auto px-5">

      <div class="flex flex-col md:flex-row md:items-end justify-between mb-12 gap-4">
        <div>
          <p class="text-xs font-bold tracking-widest text-accent uppercase mb-3" style="font-family: 'DM Mono', monospace;">
            Real stories
          </p>
          <h2 class="text-3xl md:text-4xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
            Lives changed,<br />communities built.
          </h2>
        </div>

        <div class="flex items-center gap-1">
          <Star v-for="n in 5" :key="n" :size="16" class="fill-accent text-accent" />
          <span class="text-sm text-muted-foreground ml-2" style="font-family: 'DM Mono', monospace;">
            4.9 / 5 · 6,200 reviews
          </span>
        </div>
      </div>

      <div class="grid md:grid-cols-3 gap-6">
        <div
            v-for="(t, i) in testimonials"
            :key="i"
            class="bg-card border border-border rounded-2xl p-6 hover:shadow-lg hover:shadow-black/5 transition-all"
        >
          <div class="flex items-center gap-1 mb-4">
            <Star v-for="n in t.rating" :key="n" :size="13" class="fill-accent text-accent" />
          </div>

          <p class="text-sm text-foreground leading-relaxed mb-5 italic" style="font-family: 'Fraunces', serif;">
            "{{ t.text }}"
          </p>

          <div class="flex items-center gap-3">
            <img
                :src="`https://images.unsplash.com/${t.image}?w=60&h=60&fit=crop&auto=format`"
                :alt="t.name"
                class="w-10 h-10 rounded-full object-cover border-2 border-primary/20"
            />
            <div>
              <p class="text-sm font-semibold text-foreground" style="font-family: 'Fraunces', serif;">
                {{ t.name }}
              </p>
              <p class="text-xs text-muted-foreground" style="font-family: 'DM Mono', monospace;">
                {{ t.role }}
              </p>
            </div>
          </div>
        </div>
      </div>

    </div>
  </section>
</template>


