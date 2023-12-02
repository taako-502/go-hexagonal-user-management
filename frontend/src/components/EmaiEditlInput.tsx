import { User } from '@/type/user.type'
import { Dispatch, SetStateAction } from 'react'

interface Props {
  userId: number
  editUsers: { [key: number]: User }
  setEditUsers: Dispatch<SetStateAction<{ [key: number]: User }>>
}

const EmailEditInput: React.FC<Props> = ({
  userId,
  editUsers,
  setEditUsers,
}) => {
  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const updatedUser: User = {
      ...editUsers[userId],
      email: e.target.value,
    }

    setEditUsers((prevState) => ({
      ...prevState,
      [userId]: updatedUser,
    }))
  }

  return (
    <>
      <label htmlFor={`edit_email_${userId}`} className="mx-[2em]">
        Email:
      </label>
      <input
        id={`edit_email_${userId}`}
        type="text"
        className="border bg-gray-600 border-gray-400 rounded-md p-2 mt-[8px]"
        value={editUsers[userId]?.email}
        onChange={handleEmailChange}
      />
    </>
  )
}

export default EmailEditInput
