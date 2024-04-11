package dev.openfga.language;

import dev.openfga.sdk.api.model.Userset;
import java.util.List;

public class StackRelation {
    private List<Userset> rewrites;
    private String operator;

    public StackRelation(List<Userset> rewrites, String operator) {
        this.rewrites = rewrites;
        this.operator = operator;
    }

    public List<Userset> getRewrites() {
        return rewrites;
    }

    public void setRewrites(List<Userset> rewrites) {
        this.rewrites = rewrites;
    }

    public String getOperator() {
        return operator;
    }

    public void setOperator(String operator) {
        this.operator = operator;
    }
}
