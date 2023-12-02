type Props = {
  id: number
  onClick: (id: number) => Promise<void>
}

const DeleteButton = (props: Props) => {
  return (
    <button
      className="bg-gray-500 rounded-md px-1 py-[2px] ml-[1em]"
      onClick={() => props.onClick(props.id)}
    >
      Delete
    </button>
  )
}

export default DeleteButton
