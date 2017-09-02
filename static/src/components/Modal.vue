<template>
  <div :id="id" class="modal fade" data-backdrop="static">
    <div class="modal-dialog" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">Modal title</h5>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">&times;</span>
          </button>
        </div>
        <div class="modal-body">
          <p>Modal body text goes here.</p>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-primary">Save changes</button>
          <button type="button" class="btn btn-secondary" @click="closeModal">Close</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    id: String
  },
  methods: {
    closeModal() {
      $(`#${this.id}`).modal('hide')
      this.$emit('close')
    }
  },
  mounted() {
    $(`#${this.id}`).modal('show')
    $(`#${this.id}`).attr('class', 'modal')
  }
}
</script>