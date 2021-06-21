function connectWebSocket() {
  return new Promise(function (resolve, reject) {
    let conn = new WebSocket(process.env.VUE_APP_UPBIT_URL);
    conn.binaryType = 'blob';
    conn.onopen = function () {
      resolve(conn);
    };
    conn.onerror = function (err) {
      reject(err);
    };
  });
}

export { connectWebSocket };
