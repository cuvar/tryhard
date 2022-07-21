<script>
  import { onMount } from 'svelte';
  import { createScene } from "./scene.js";
  /**
* @type {HTMLCanvasElement}
*/
  let el;
  onMount(() => {
    createScene(el)
  });
</script>

<canvas bind:this={el}></canvas>