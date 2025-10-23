import { useEffect, useState } from "react";
import { getProducts, createProduct, updateProduct, deleteProduct } from "./api";

function App() {
  const [products, setProducts] = useState([]);
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [price, setPrice] = useState("");
  const [editingId, setEditingId] = useState(null);

  useEffect(() => {
    getProducts().then(setProducts);
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const productData = { name, description, price: parseFloat(price) };

    if (editingId) {
      const updated = await updateProduct(editingId, productData);
      setProducts(products.map(p => p.id === editingId ? updated : p));
      setEditingId(null);
    } else {
      const newProduct = await createProduct(productData);
      setProducts([...products, newProduct]);
    }

    setName("");
    setDescription("");
    setPrice("");
  };

  const handleEdit = (product) => {
    setName(product.name);
    setDescription(product.description);
    setPrice(product.price);
    setEditingId(product.id);
  };

  const handleDelete = async (id) => {
    await deleteProduct(id);
    setProducts(products.filter(p => p.id !== id));
  };

  return (
    <div style={{ padding: "2rem" }}>
      <h1>Hola, soy Jhonatan Zuñiga 👋</h1>
      <h2>Productos</h2>
      <ul>
        {products.map((p) => (
          <li key={p.id}>
            <strong>{p.name}</strong> - ${p.price} <br />
            Descripción: {p.description} <br />
            <button onClick={() => handleEdit(p)}>Editar</button>
            <button onClick={() => handleDelete(p.id)}>Eliminar</button>
          </li>
        ))}
      </ul>
      <form onSubmit={handleSubmit}>
        <input
          placeholder="Nombre"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <input
          placeholder="Descripción"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <input
          placeholder="Precio"
          type="number"
          step="0.01"
          value={price}
          onChange={(e) => setPrice(e.target.value)}
          required
        />
        <button type="submit">{editingId ? "Actualizar" : "Agregar"}</button>
      </form>
    </div>
  );
}

export default App;
