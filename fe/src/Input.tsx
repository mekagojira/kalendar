export const Input = (props: {
    label: string
    value: string
    id: string
    onChange: (value: string) => void
}) => {
    return (
        <div className="flex flex-col">
            <label htmlFor={props.id} className="font-bold">
                {props.label}
            </label>
            <input
                id={props.id}
                type="text"
                value={props.value}
                className="px-3 py-2 border border-purple-200"
                onChange={(e) => props.onChange(e.target.value)}
            />
        </div>
    )
}
