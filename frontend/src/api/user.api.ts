import { User, isUserType, isUsersType } from '@/type/user.type'
import { AxiosInstance } from 'axios'

class UserRepository {
  private axios: AxiosInstance

  constructor(axiosInstance: AxiosInstance) {
    this.axios = axiosInstance
  }

  findAll = async (): Promise<User[]> => {
    const response = await this.axios.get('/users')
    if (isUsersType(response.data)) return response.data
    throw Error('API Error')
  }

  create = async (user: User): Promise<User> => {
    const response = await this.axios.post('/user', user)
    if (isUserType(response.data)) return response.data
    throw Error('API Error')
  }

  update = async (user: User) => {
    await this.axios.put(`/user/${user.id}`, user)
  }

  delete = async (id: number) => {
    await this.axios.delete(`/user/${id}`)
  }
}
export default UserRepository
