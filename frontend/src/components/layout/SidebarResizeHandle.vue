<template>
  <div
    class="sidebar-resize-handle hidden lg:block"
    role="separator"
    tabindex="0"
    aria-orientation="vertical"
    :aria-label="label"
    :aria-valuemin="SIDEBAR_MIN_WIDTH"
    :aria-valuemax="SIDEBAR_MAX_WIDTH"
    :aria-valuenow="width"
    @pointerdown.prevent="startResize"
    @keydown="resizeWithKeyboard"
  />
</template>

<script setup lang="ts">
import { onBeforeUnmount } from 'vue'
import {
  clampSidebarWidth,
  SIDEBAR_MIN_WIDTH,
  SIDEBAR_MAX_WIDTH
} from '@/utils/sidebarWidth'

const props = defineProps<{ width: number; label: string }>()
const emit = defineEmits<{ resize: [width: number] }>()
let cleanup: (() => void) | undefined

function startResize(event: PointerEvent) {
  if (event.button !== 0) return
  cleanup?.()
  // pointerdown.prevent suppresses the browser's default focus behavior.
  const handle = event.currentTarget as HTMLElement
  handle.focus()
  const startX = event.clientX
  const startWidth = props.width
  const pointerId = event.pointerId
  const body = document.body
  const previousCursor = body.style.cursor
  const previousUserSelect = body.style.userSelect
  body.classList.add('is-resizing')
  body.style.cursor = 'col-resize'
  body.style.userSelect = 'none'

  const move = (next: PointerEvent) => {
    if (next.pointerId !== pointerId) return
    emit('resize', clampSidebarWidth(startWidth + next.clientX - startX))
  }
  const finish = () => cleanup?.()
  const release = (next: PointerEvent) => {
    if (next.pointerId === pointerId) finish()
  }
  cleanup = () => {
    window.removeEventListener('pointermove', move)
    window.removeEventListener('pointerup', release)
    window.removeEventListener('pointercancel', release)
    window.removeEventListener('blur', finish)
    body.classList.remove('is-resizing')
    body.style.cursor = previousCursor
    body.style.userSelect = previousUserSelect
    cleanup = undefined
  }
  window.addEventListener('pointermove', move)
  window.addEventListener('pointerup', release)
  window.addEventListener('pointercancel', release)
  window.addEventListener('blur', finish)
}

function resizeWithKeyboard(event: KeyboardEvent) {
  const widths: Record<string, number> = {
    ArrowLeft: props.width - 10,
    ArrowRight: props.width + 10,
    Home: SIDEBAR_MIN_WIDTH,
    End: SIDEBAR_MAX_WIDTH
  }
  const width = widths[event.key]
  if (width === undefined) return
  event.preventDefault()
  emit('resize', clampSidebarWidth(width))
}

onBeforeUnmount(() => cleanup?.())
</script>

<style scoped>
.sidebar-resize-handle {
  position: absolute;
  top: 0;
  right: -3px;
  width: 6px;
  height: 100%;
  cursor: col-resize;
  touch-action: none;
  z-index: 1;
  transition: background-color 0.15s ease;
}

.sidebar-resize-handle:hover,
.sidebar-resize-handle:focus-visible,
.sidebar-resize-handle:active {
  background-color: rgba(0, 122, 255, 0.3);
}

@media (prefers-reduced-motion: reduce) {
  .sidebar-resize-handle {
    transition: none;
  }
}
</style>
