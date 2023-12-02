import UserRepository from '@/api/user.api'
import axios from 'axios'

const useRepository = () => {
  const instance = axios.create({
    baseURL: process.env.NEXT_PUBLIC_BASE_URL,
  })
  const repository = {
    user: new UserRepository(instance),
  }
  return repository
}

export default useRepository
