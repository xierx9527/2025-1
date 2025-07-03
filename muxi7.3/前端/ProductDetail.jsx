import { useParams } from 'react-router-dom';

export default function ProductDetail() {
    const { id } = useParams();

    return (
        <div>
            <h2>商品详情</h2>
            <p>当前查看的商品ID: {id}</p>
        </div>
    );
}