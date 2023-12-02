import useRepository from '@/hooks/useAxios'
import useAxios from '@/hooks/useAxios'
import { User } from '@/type/user.type'
import axios from 'axios'
import { Fragment, useEffect, useState } from 'react'

export default function Home() {
  const [newUsers, setNewUsers] = useState<User[]>([])
  const [editUsers, setEditUsers] = useState<{ [key: number]: User }>({})
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [error, setError] = useState('')
  const [editMode, setEditMode] = useState<Map<number, boolean>>(new Map())
  const repository = useRepository()

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const users = await repository.user.findAll()
        setNewUsers(users)
        const editUsers: { [key: number]: User } = users.reduce(
          (acc, user) => ({
            ...acc,
            [user.id as number]: user,
          }),
          {},
        )
        setEditUsers(editUsers)
      } catch (error) {
        console.error(error)
        if (axios.isAxiosError(error) && error.response) {
          // 404エラー以外はコンソールに出力する
          if (error.response.status !== 404) {
            console.error(error.message)
            setError(error.message)
          }
        }
      }
    }
    fetchUsers()
  }, [])

  const newUser = async () => {
    const body: User = { username, email }
    try {
      await repository.user.create(body)
      setUsername('')
      setEmail('')
      setError('')
      // 一旦画面をリロードする
      window.location.reload()
    } catch (error) {
      console.error(error)
      if (error instanceof Error) {
        console.error(error.message)
        setError(error.message)
      }
    }
  }

  const update = async (id: number) => {
    const user: User = {
      id,
      username: editUsers[id].username || '',
      email: editUsers[id].email || '',
    }
    try {
      await repository.user.update(user)
      setUsername('')
      setEmail('')
      setError('')
      // 一旦画面をリロードする
      window.location.reload()
    } catch (error) {
      console.error(error)
      if (axios.isAxiosError(error) && error.response) {
        console.error(error.message)
        setError(error.message)
      }
    }
  }

  const userDelete = async (id: number) => {
    try {
      await repository.user.delete(id)
      // 一旦画面をリロードする
      window.location.reload()
    } catch (error) {
      console.error(error)
      if (axios.isAxiosError(error) && error.response) {
        console.error(error.message)
        setError(error.message)
      }
    }
  }

  const handleSubmit = (event: any) => {
    event.preventDefault()
  }

  const toggleEditMode = (id: number) => {
    setEditMode((prev) => {
      const newEditMode = new Map(prev)
      newEditMode.set(id, !newEditMode.get(id))
      return newEditMode
    })
  }

  return (
    <main className="flex min-h-screen flex-col items-center p-24">
      <h1 className="text-2xl">ユーザ管理画面</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="username" className="mr-[2em]">
            ユーザ名:
          </label>
          <input
            id="username"
            type="text"
            className="border bg-gray-600 border-gray-400 rounded-md p-2 mt-[8px]"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>
        <div>
          <label htmlFor="email" className="mr-[2em]">
            メールアドレス:
          </label>
          <input
            id="email"
            type="email"
            className="border bg-gray-600 border-gray-400 rounded-md p-2 mt-[8px]"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <div className="text-red-500">{error}</div>
        <button
          className="bg-blue-500 text-white cursor-pointer rounded-md p-2 mt-[1em]"
          onClick={newUser}
        >
          save
        </button>
      </form>
      <ul>
        {newUsers.map((user) => (
          <li key={user.id}>
            <pre className="mt-1">
              {editMode.get(user.id as number) ? (
                <span>
                  <label htmlFor="edit_username" className="mr-[2em]">
                    Name:
                  </label>
                  <input
                    id="edit_username"
                    type="text"
                    className="border bg-gray-600 border-gray-400 rounded-md p-2 mt-[8px]"
                    value={editUsers[user.id as number]?.username || ''}
                    onChange={(e) =>
                      setEditUsers({
                        ...editUsers,
                        [user.id as number]: {
                          ...editUsers[user.id as number],
                          username: e.target.value,
                        },
                      })
                    }
                  />
                  <label htmlFor="edit_email" className="mr-[2em]">
                    Email:
                  </label>
                  <input
                    id="edit_email"
                    type="text"
                    className="border bg-gray-600 border-gray-400 rounded-md p-2 mt-[8px]"
                    value={editUsers[user.id as number]?.email || ''}
                    onChange={(e) =>
                      setEditUsers({
                        ...editUsers,
                        [user.id as number]: {
                          ...editUsers[user.id as number],
                          email: e.target.value,
                        },
                      })
                    }
                  />
                </span>
              ) : (
                <Fragment>
                  Name: {user.username}&#009;&#009;Email: {user.email}
                </Fragment>
              )}

              <button
                className="bg-gray-500 rounded-md px-1 py-[2px] ml-[1em]"
                onClick={() => toggleEditMode(user.id as number)}
              >
                {editMode.get(user.id as number) ? 'Cancel' : 'Edit'}
              </button>
              {editMode.get(user.id as number) && (
                <button
                  className="bg-gray-500 rounded-md px-1 py-[2px] ml-[1em]"
                  onClick={() => update(user.id as number)}
                >
                  Update
                </button>
              )}
              <button
                className="bg-gray-500 rounded-md px-1 py-[2px] ml-[1em]"
                onClick={() => {
                  userDelete(user.id as number)
                }}
              >
                Delete
              </button>
            </pre>
          </li>
        ))}
      </ul>
    </main>
  )
}
