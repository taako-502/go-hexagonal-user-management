export interface User {
  id?: number
  username: string
  email: string
}

export const isUserType = (arg: any): arg is User => {
  return typeof arg.username === 'string' && typeof arg.email === 'string'
}

export const isUsersType = (arg: any): arg is User[] => {
  return Array.isArray(arg) && arg.every((e) => isUserType(e))
}
