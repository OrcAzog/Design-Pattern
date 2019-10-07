import color.Color;
import shape.Shape;

/**
 * @program: Solution
 * @description: ${DESCRIPTION}
 * @author: 香喷喷大猪蹄子
 * @create: 2019-10-07 13:42
 **/
public abstract class AbstractFactory  {
    public abstract Shape getShape();
    public abstract Color getColor();
}
