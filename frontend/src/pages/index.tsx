import EmailInput from '@/components/inputs/EmailInput'
import NameInput from '@/components/inputs/NameInput'
import SaveButton from '@/components/buttons/SaveButton'
import UserList from '@/components/UserList'
import useRepository from '@/hooks/useRepository'
import { User } from '@/type/user.type'
import axios from 'axios'
import { useState } from 'react'

export default function Home() {
  const [users, setUsers] = useState<{ [key: number]: User }>({})
  const [editUsers, setEditUsers] = useState<{ [key: number]: User }>({})
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [error, setError] = useState('')
  const repository = useRepository()

  const newUser = async () => {
    const body: User = { username, email }
    try {
      const user = await repository.user.create(body)
      clear()

      // 新規ユーザを追加
      const addNewUser = (
        prevState: {
          [key: number]: User
        },
        newUser: User,
      ) => ({
        ...prevState,
        [newUser.id as number]: newUser,
      })

      setUsers((prev) => addNewUser(prev, user))
      setEditUsers((prev) => addNewUser(prev, user))
    } catch (error) {
      if (axios.isAxiosError(error) && error.response) {
        console.error(error.message)
        setError(error.response.data.message)
      }
    }
  }

  const handleSubmit = (event: any) => {
    event.preventDefault()
  }

  const clear = () => {
    setUsername('')
    setEmail('')
    setError('')
  }

  return (
    <main className="flex min-h-screen flex-col items-center p-24">
      <h1 className="text-2xl">ユーザ管理画面</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <NameInput username={username} setUsername={setUsername} />
        </div>
        <div>
          <EmailInput email={email} setEmail={setEmail} />
        </div>
        <div className="text-red-500">{error}</div>
        <SaveButton onClick={newUser} />
      </form>
      <UserList
        clear={clear}
        users={users}
        setUsers={setUsers}
        editUsers={editUsers}
        setEditUsers={setEditUsers}
        setError={setError}
      />
    </main>
  )
}
