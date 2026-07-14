<script setup>
import { computed } from 'vue'
import { Clock, MapPin } from '@lucide/vue'

const props = defineProps({
  campaign: {
    type: Object,
    required: true
  }
})

const pct = computed(() => {
  return Math.round((props.campaign.raised / props.campaign.goal) * 100)
})
</script>

<template>
  <div class="bg-card rounded-2xl border border-border overflow-hidden hover:shadow-xl hover:shadow-black/5 transition-all duration-300 hover:-translate-y-0.5 group flex flex-col">

    <div class="relative overflow-hidden">
      <img
          :src="`https://images.unsplash.com/${campaign.image}?w=500&h=260&fit=crop&auto=format`"
          :alt="campaign.title"
          class="w-full h-44 object-cover group-hover:scale-105 transition-transform duration-500"
      />

      <div class="absolute top-3 left-3">
        <span class="bg-white/90 backdrop-blur-sm text-foreground text-xs font-semibold px-2.5 py-1 rounded-full" style="font-family: 'DM Mono', monospace;">
          {{ campaign.category }}
        </span>
      </div>

      <div class="absolute top-3 right-3 flex items-center gap-1.5 bg-white/90 backdrop-blur-sm text-foreground text-xs font-semibold px-2.5 py-1 rounded-full">
        <Clock :size="11" />
        {{ campaign.daysLeft }}d left
      </div>
    </div>

    <div class="p-5 flex flex-col flex-1">

      <div class="flex items-center gap-2 mb-3">
        <div
            class="w-7 h-7 rounded-full flex items-center justify-center text-white text-xs font-bold flex-shrink-0"
            :style="{ background: campaign.avatarColor, fontFamily: 'Fraunces, serif' }"
        >
          {{ campaign.avatar }}
        </div>
        <div class="flex items-center gap-1 text-xs text-muted-foreground" style="font-family: 'DM Mono', monospace;">
          <MapPin :size="10" />
          {{ campaign.location }}
        </div>
      </div>

      <h3 class="font-bold text-foreground text-base mb-2 leading-snug flex-1" style="font-family: 'Fraunces', serif;">
        {{ campaign.title }}
      </h3>
      <p class="text-sm text-muted-foreground mb-4 leading-relaxed" style="font-family: 'DM Sans', sans-serif;">
        {{ campaign.description }}
      </p>

      <div class="mt-auto">
        <div class="flex justify-between items-center mb-1.5">
          <span class="text-sm font-bold text-foreground" style="font-family: 'DM Mono', monospace;">
            ${{ campaign.raised.toLocaleString() }}
          </span>
          <span class="text-xs text-muted-foreground" style="font-family: 'DM Mono', monospace;">
            {{ pct }}% of ${{ campaign.goal.toLocaleString() }}
          </span>
        </div>

        <div class="w-full h-2 bg-muted rounded-full overflow-hidden mb-3">
          <div
              class="h-full bg-primary rounded-full"
              :style="{ width: `${pct}%`, transition: 'width 0.6s ease' }"
          ></div>
        </div>

        <div class="flex items-center justify-between">
          <span class="text-xs text-muted-foreground" style="font-family: 'DM Mono', monospace;">
            {{ campaign.backers }} backers
          </span>
          <button class="text-xs font-semibold bg-primary text-primary-foreground px-3.5 py-1.5 rounded-lg hover:opacity-90 transition-opacity">
            Contribute
          </button>
        </div>
      </div>

    </div>
  </div>
</template>

