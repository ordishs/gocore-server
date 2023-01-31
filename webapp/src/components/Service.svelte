<script>
  import { beforeUpdate } from 'svelte'
  import { humanTime, filterJSON } from '../lib/utils'

  export let json = {}
  export let details = false

  export let removeMe

  let host = ''
  let address = ''
  let serviceName = ''
  let uptime = ''
  let lastHeartbeat = ''
  let colour = 'panel-danger'
  let removeButton = false

  beforeUpdate(function() {
    host = json.host
    address = json.address
    serviceName = json.serviceName
    uptime = humanTime(
      new Date().getTime() - new Date(json.startTime).getTime()
    )

    const lastHeartbeatMillis =
      new Date().getTime() - new Date(json._ts).getTime()
    if (lastHeartbeatMillis < 45 * 1000) {
      colour = 'panel-success'
      removeButton = false
    } else if (lastHeartbeatMillis < 60 * 60 * 1000) {
      colour = 'panel-warning'
      removeButton = false
    } else {
      colour = 'panel-danger'
      removeButton = true
    }

    lastHeartbeat = humanTime(lastHeartbeatMillis)
  })
</script>

<style>
  pre {
    max-height: 120px;
    overflow-y: scroll;
    background-color: white;
    border: none;
    padding: 0;
  }

  button {
    background: none;
    border: none;
    position: absolute;
    top: 0;
    right: 0;
    margin: 4px;
  }

  .panel {
    margin-bottom: 0;
  }
</style>

<div class="col-lg-3 col-md-4 col-sm-6 col-xs-12" style="padding: 5px;">
  <div class="panel {colour}">
    <div class="panel-heading">
      {#if removeButton}
        <button on:click={removeMe(json)}>
          <span class="glyphicon glyphicon-remove" />
        </button>
      {/if}

      <div class="row">
        <div class="checkbox">
          <form>
          <label for={'div'}>
            {serviceName}
            <div style="font-size: 0.8em;">
              <div>Listening on {address}</div>
              <div>{host}</div>
              <div>Up {uptime}</div>
              <div>Last heatbeat {lastHeartbeat} ago</div>
            </div>
          </label>
        </div>
      </div>
    </div>
    {#if details}
      <div class="panel-body">
        <pre style="font-size: 0.7em;">
          {JSON.stringify(filterJSON(json), null, 2)}
        </pre>
      </div>
    {/if}
  </div>
</div>
