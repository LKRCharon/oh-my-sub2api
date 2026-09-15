import { afterEach, describe, expect, it } from 'vitest'
import { mount, type VueWrapper } from '@vue/test-utils'
import SidebarResizeHandle from '../SidebarResizeHandle.vue'

let wrapper: VueWrapper | undefined

function pointer(target: EventTarget, type: string, clientX: number, pointerId = 1) {
  target.dispatchEvent(Object.assign(new Event(type, { bubbles: true, cancelable: true }), {
    clientX, pointerId, button: 0
  }))
}

function createHandle() {
  wrapper = mount(SidebarResizeHandle, {
    attachTo: document.body,
    props: { width: 256, label: 'Resize sidebar' }
  })
  return wrapper
}

afterEach(() => {
  wrapper?.unmount()
  wrapper = undefined
  document.body.style.cursor = ''
  document.body.style.userSelect = ''
})

describe('SidebarResizeHandle', () => {
  it('drags from the current width and clamps both ends', () => {
    const handle = createHandle()
    pointer(handle.element, 'pointerdown', 256)
    expect(document.activeElement).toBe(handle.element)
    pointer(window, 'pointermove', 316)
    pointer(window, 'pointermove', 900)
    pointer(window, 'pointermove', -100)
    expect(handle.emitted('resize')).toEqual([[316], [400], [180]])
    pointer(window, 'pointerup', -100)
    pointer(window, 'pointermove', 300)
    expect(handle.emitted('resize')).toHaveLength(3)
  })

  it('ignores movement and release from a different pointer', () => {
    const handle = createHandle()
    pointer(handle.element, 'pointerdown', 256)
    pointer(window, 'pointermove', 400, 2)
    pointer(window, 'pointerup', 400, 2)
    expect(handle.emitted('resize')).toBeUndefined()
    pointer(window, 'pointermove', 300)
    expect(handle.emitted('resize')).toEqual([[300]])
  })

  it.each(['pointerup', 'pointercancel', 'blur', 'unmount'])('restores page interaction after %s', (finish) => {
    const handle = createHandle()
    document.body.style.cursor = 'crosshair'
    document.body.style.userSelect = 'text'
    pointer(handle.element, 'pointerdown', 256)
    expect(document.body.classList.contains('is-resizing')).toBe(true)
    if (finish === 'unmount') {
      handle.unmount()
      wrapper = undefined
    } else if (finish === 'blur') {
      window.dispatchEvent(new Event('blur'))
    } else {
      pointer(window, finish, 256)
    }
    expect(document.body.classList.contains('is-resizing')).toBe(false)
    expect(document.body.style.cursor).toBe('crosshair')
    expect(document.body.style.userSelect).toBe('text')
    pointer(window, 'pointermove', 360)
    expect(handle.emitted('resize')).toBeUndefined()
  })

  it('supports keyboard resizing and reports the current accessible value', async () => {
    const handle = createHandle()
    expect(handle.attributes('role')).toBe('separator')
    expect(handle.attributes('aria-label')).toBe('Resize sidebar')
    await handle.trigger('keydown', { key: 'ArrowRight' })
    await handle.setProps({ width: 266 })
    expect(handle.attributes('aria-valuenow')).toBe('266')
    await handle.trigger('keydown', { key: 'ArrowLeft' })
    await handle.trigger('keydown', { key: 'Home' })
    await handle.trigger('keydown', { key: 'End' })
    await handle.trigger('keydown', { key: 'Tab' })
    expect(handle.emitted('resize')).toEqual([[266], [256], [180], [400]])
  })
})
