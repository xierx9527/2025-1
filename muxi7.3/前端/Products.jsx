import { Link } from 'react-router-dom';

export default function Products() {
    const products = [
        { id: 1, name: '商品1' },
        { id: 2, name: '商品2' },
        { id: 3, name: '商品3' },
    ];

    return (
        <div>
            <h2>商品列表</h2>
            <ul>
                {products.map(product => (
                    <li key={product.id}>
                        <Link to={`/products/${product.id}`}>{product.name}</Link>
                    </li>
                ))}
            </ul>
        </div>
    );
}