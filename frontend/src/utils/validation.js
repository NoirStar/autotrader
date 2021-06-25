function validateEmail(email) {
  const re =
    /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  return re.test(String(email).toLowerCase());
}

function validatePassword(pw) {
  const symbol = /[`~!@#$%^&*\\'";:?]/;
  const eng = /[a-zA-Z]/;
  const num = /[0-9]/;

  return symbol.test(pw) && eng.test(pw) && num.test(pw);
}

export { validateEmail, validatePassword };
