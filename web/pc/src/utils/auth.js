
const TokenKey = 'Admin-Token'
const MenuType = 'Menu-Type'
export function getToken() {
  return  localStorage.getItem(TokenKey)
}
export function setToken(token) {
  return  localStorage.setItem(TokenKey,token)
}
export function removeToken() {
  removeMenuType()
  return  localStorage.removeItem(TokenKey)
}


export function setMenuType(ty){
  return  localStorage.setItem(MenuType,ty)
}
export function getMenuType() {
  return  localStorage.getItem(MenuType)
}
export function removeMenuType() {
  return  localStorage.removeItem(MenuType)
}
