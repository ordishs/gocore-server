function humanTime (timeInMillis) {
  timeInMillis = timeInMillis / 1000
  const days = Math.floor(timeInMillis / 86400)
  const hours = Math.floor((timeInMillis % 86400) / 3600)
  const minutes = Math.floor(((timeInMillis % 86400) % 3600) / 60)
  const seconds = Math.floor((((timeInMillis % 86400) % 3600) % 60))

  let difference = ''

  if (days > 0) {
    difference += days
    difference += ' day' + (days === 1 ? ', ' : 's, ')
    difference += hours
    difference += ' hour' + (hours === 1 ? ', ' : 's, ')
    difference += minutes
    difference += ' minute' + (minutes === 1 ? ' and ' : 's and ')
    difference += seconds
    difference += ' second' + (seconds === 1 ? '' : 's')
    return difference
  }

  if (hours > 0) {
    difference += hours
    difference += ' hour' + (hours === 1 ? ', ' : 's, ')
    difference += minutes
    difference += ' minute' + (minutes === 1 ? ' and ' : 's and ')
    difference += seconds
    difference += ' second' + (seconds === 1 ? '' : 's')
    return difference
  }

  if (minutes > 0) {
    difference += minutes
    difference += ' minute' + (minutes === 1 ? ' and ' : 's and ')
    difference += seconds
    difference += ' second' + (seconds === 1 ? '' : 's')
    return difference
  }

  if (seconds > 0) {
    difference += seconds
    difference += ' second' + (seconds === 1 ? '' : 's')
    return difference
  }

  return '0 seconds'
}

function filterJSON (json) {
  const result = {}

  for (const key in json) {
    switch (key) {
      case 'host':
      case 'address':
      case 'serviceName':
      case '_ts':
        // Don't add
        break
      default:
        result[key] = json[key]
    }
  }

  return result
};

module.exports = {
  humanTime,
  filterJSON
}
