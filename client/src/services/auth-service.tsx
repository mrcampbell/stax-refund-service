export const GetLocalToken = () => {
    return localStorage.getItem('token')
}

export const SetLocalToken = (token: string) => {
    localStorage.setItem('token', token)
}

export const RemoveLocalToken = () => {
    localStorage.removeItem('token')
}