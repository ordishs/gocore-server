<script>
  export let filter = ''
  export let details = false
</script>

<style>
  label {
    color: #9d9d9d;
    font-weight: normal;
    margin-left: 20px;
    margin-right: 10px;
  }

  input {
    autocapitalize: off;
    autocorrect: off;
  }
</style>

<nav class="navbar navbar-inverse navbar-fixed-top">
  <div class="container-fluid">
    <div class="navbar-header">
      <span class="navbar-brand">GoCore Monitor</span>
    </div>

    <div class="navbar-collapse">
      <!-- Search -->
      <form class="navbar-form navbar-right" role="search">
        <div class="form-group">
          <label class="checkbox">
            <input
              type="checkbox"
              class="form-console"
              bind:checked={details} />
            Details
          </label>

          <input
            type="search"
            placeholder="Filter tokens..."
            class="form-control"
            bind:value={filter} />
        </div>
      </form>
    </div>
  </div>
</nav>
