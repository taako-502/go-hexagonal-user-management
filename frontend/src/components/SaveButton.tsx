type Props = {
  onClick: () => Promise<void>
}

const SaveButton = (props: Props) => {
  return (
    <button
      className="bg-blue-500 text-white cursor-pointer rounded-md p-2 mt-[1em]"
      onClick={props.onClick}
    >
      save
    </button>
  )
}

export default SaveButton
