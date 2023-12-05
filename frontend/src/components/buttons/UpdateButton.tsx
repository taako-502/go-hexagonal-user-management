type Props = {
  id: number
  onClick: (id: number) => Promise<void>
}

const UpdateButton: React.FC<Props> = ({ id, onClick }) => {
  return (
    <button
      className="bg-gray-500 rounded-md px-1 py-[2px] ml-[1em]"
      onClick={() => onClick(id)}
    >
      Update
    </button>
  )
}

export default UpdateButton
