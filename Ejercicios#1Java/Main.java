import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

// Clase Persona
class Persona {
    private String nombre;
    private String apellido;
    private String telefono;

    public Persona(String nombre, String apellido, String telefono) {
        this.nombre = nombre;
        this.apellido = apellido;
        this.telefono = telefono;
    }

    @Override
    public String toString() {
        return "Persona: " + nombre + " " + apellido + ", Teléfono: " + telefono;
    }
}

// Clase Contacto
class Contacto extends Persona {
    public Contacto(String nombre, String apellido, String telefono) {
        super(nombre, apellido, telefono);
    }

    @Override
    public String toString() {
        return "Contacto -> " + super.toString();
    }
}

// Clase ContactoFamiliar
class ContactoFamiliar extends Contacto {
    private String parentesco;

    public ContactoFamiliar(String nombre, String apellido, String telefono, String parentesco) {
        super(nombre, apellido, telefono);
        this.parentesco = parentesco;
    }

    @Override
    public String toString() {
        return "Contacto Familiar -> " + super.toString() + ", Parentesco: " + parentesco;
    }
}

// Clase ContactoEmpresarial
class ContactoEmpresarial extends Contacto {
    private String empresa;
    private String puesto;

    public ContactoEmpresarial(String nombre, String apellido, String telefono, String empresa, String puesto) {
        super(nombre, apellido, telefono);
        this.empresa = empresa;
        this.puesto = puesto;
    }

    @Override
    public String toString() {
        return "Contacto Empresarial -> " + super.toString() + ", Empresa: " + empresa + ", Puesto: " + puesto;
    }
}

// Clase Evento
class Evento {
    private String nombreEvento;
    private String fecha;

    public Evento(String nombreEvento, String fecha) {
        this.nombreEvento = nombreEvento;
        this.fecha = fecha;
    }

    @Override
    public String toString() {
        return "Evento: " + nombreEvento + ", Fecha: " + fecha;
    }
}

// Clase EventoEspecifico
class EventoEspecifico extends Evento {
    private String descripcion;

    public EventoEspecifico(String nombreEvento, String fecha, String descripcion) {
        super(nombreEvento, fecha);
        this.descripcion = descripcion;
    }

    @Override
    public String toString() {
        return "Evento Específico -> " + super.toString() + ", Descripción: " + descripcion;
    }
}

// Singleton para Agenda
class Agenda {
    private static Agenda instancia;
    private List<Object> elementos;

    private Agenda() {
        elementos = new ArrayList<>();
    }

    public static Agenda getInstancia() {
        if (instancia == null) {
            instancia = new Agenda();
        }
        return instancia;
    }

    public void agregarElemento(Object elemento) {
        elementos.add(elemento);
    }

    public void eliminarElemento(Object elemento) {
        elementos.remove(elemento);
    }

    public void modificarElemento(int index, Object nuevoElemento) {
        elementos.set(index, nuevoElemento);
    }

    public List<Object> getElementos() {
        return elementos;
    }

    @Override
    public String toString() {
        return "Agenda: " + elementos;
    }
}

// Abstract Factory para Contactos y Eventos
abstract class AbstractFactory {
    abstract Contacto crearContacto(String nombre, String apellido, String telefono, String tipo, String... extra);
    abstract Evento crearEvento(String nombreEvento, String fecha, String tipo, String... extra);
}

class ContactoFactory extends AbstractFactory {
    @Override
    public Contacto crearContacto(String nombre, String apellido, String telefono, String tipo, String... extra) {
        if (tipo.equalsIgnoreCase("familiar")) {
            return new ContactoFamiliar(nombre, apellido, telefono, extra[0]);
        } else if (tipo.equalsIgnoreCase("empresarial")) {
            return new ContactoEmpresarial(nombre, apellido, telefono, extra[0], extra[1]);
        }
        return new Contacto(nombre, apellido, telefono);
    }

    @Override
    public Evento crearEvento(String nombreEvento, String fecha, String tipo, String... extra) {
        return null;
    }
}

class EventoFactory extends AbstractFactory {
    @Override
    public Contacto crearContacto(String nombre, String apellido, String telefono, String tipo, String... extra) {
        return null;
    }

    @Override
    public Evento crearEvento(String nombreEvento, String fecha, String tipo, String... extra) {
        if (tipo.equalsIgnoreCase("especifico")) {
            return new EventoEspecifico(nombreEvento, fecha, extra[0]);
        }
        return new Evento(nombreEvento, fecha);
    }
}

// Clase principal
public class Main {
    public static void main(String[] args) {
        // Singleton de la Agenda
        Agenda agenda = Agenda.getInstancia();

        // Factories
        AbstractFactory contactoFactory = new ContactoFactory();
        AbstractFactory eventoFactory = new EventoFactory();

        // Crear contactos y eventos
        Contacto c1 = contactoFactory.crearContacto("Juan", "Pérez", "123456", "familiar", "Hermano");
        Contacto c2 = contactoFactory.crearContacto("Ana", "Gómez", "654321", "empresarial", "TechCorp", "Ingeniera");
        Evento e1 = eventoFactory.crearEvento("Reunión", "2024-06-01", "simple");
        Evento e2 = eventoFactory.crearEvento("Conferencia", "2024-07-15", "especifico", "Tecnología e Innovación");

        // Agregar a la agenda
        agenda.agregarElemento(c1);
        agenda.agregarElemento(e1);
        agenda.agregarElemento(c2);
        agenda.agregarElemento(e2);

        // Imprimir elementos de la agenda
        System.out.println("Todos los elementos en la agenda:");
        System.out.println(agenda);

        // Filtrar contactos y eventos
        List<Object> contactos = agenda.getElementos().stream()
                .filter(elemento -> elemento instanceof Contacto)
                .collect(Collectors.toList());
        List<Object> eventos = agenda.getElementos().stream()
                .filter(elemento -> elemento instanceof Evento)
                .collect(Collectors.toList());

        System.out.println("\nContactos en la agenda:");
        contactos.forEach(System.out::println);

        System.out.println("\nEventos en la agenda:");
        eventos.forEach(System.out::println);
    }
}
